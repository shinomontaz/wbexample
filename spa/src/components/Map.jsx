
import React, { useRef, useState, useContext, useEffect } from "react"
import * as ol from "ol";
import TileLayer from "ol/layer/Tile";
import VectorLayer from 'ol/layer/Vector';
import VectorSource from 'ol/source/Vector';
import Feature from 'ol/Feature';
import { Circle, Fill, Stroke, Text, Style, Icon } from "ol/style";

import {Point, Polygon, LineString} from 'ol/geom';
import {defaults as defaultControls} from 'ol/control.js';
import OSMSource from "ol/source/OSM";
import { useMapContext } from "../contexts/MapContext";

import {unByKey} from "ol/Observable";

import { transform, fromLonLat, toLonLat, transformExtent } from "ol/proj";
import {toStringHDMS} from 'ol/coordinate.js';

import style from "../styles";
import api from '../api';

export default function Map({ children, zoom, center, mapMode }) {
    const mapRef = useRef();

    const { map, setMap, sourceWh, setWhSource, sourceCar, setCarSource, sourceLink, setLinkSource, setViewport, points, trucks } = useMapContext();
    const [currClickId, setClickId ] = useState(null);

    useEffect(() => {
      let sWhVec = new VectorSource();
      let sCarVec = new VectorSource();
      let sLinkVec = new VectorSource();

   		let options = {
   			view: new ol.View({ zoom, center }),
   			layers: [
          new TileLayer({
            source: new OSMSource()
          }),
          new VectorLayer({
            source: sWhVec,
            style: (feature) => {
              return [
                new Style({
                image: new Circle({
                radius: 10,
                fill: new Fill({
                  //af52de
                    color: '#af52de',
                }),
                stroke: new Stroke({
                    width: 1,
                    color: '#000'
                }),
                })
              }),
                new Style({
                image: new Icon({
    src: '/icons/wh.png',
    scale: 0.025
  })
              })
            ]
            }
          }),
          new VectorLayer({
            source: sLinkVec,
            style: new Style({
              fill: new Fill({ color: '#00FF00', weight: 4 }),
              stroke: new Stroke({
        width: 3,
        color: 'rgba(100, 100, 255, 1)',
        lineDash: [.1, 5] //or other combinations
    }),
            })
          }),
          new VectorLayer({
            source: sCarVec,
            style: (feature) => {
              return [
                new Style({
                image: new Circle({
                radius: 10,
                fill: new Fill({
                    color: '#fff',
                  }),
                stroke: new Stroke({
                    width: 1,
                    color: '#af52de'
                  }),
                })
              }),
              new Style({
                image: new Icon({
                src: '/icons/van.svg',
                scale: 0.02
                })
              })
            ];
            }
          })
        ],
   			controls: [],
   			overlays: []
   		};

   		let mapObject = new ol.Map(options);

   		mapObject.setTarget(mapRef.current);

      var viewport = transformExtent( mapObject.getView().calculateExtent(mapObject.getSize()), 'EPSG:3857', 'EPSG:4326' );
      setViewport(viewport);

      setWhSource(sWhVec);
      setCarSource(sCarVec);
      setLinkSource(sLinkVec);

      setMap(mapObject);

   		return () => mapObject.setTarget(undefined);
   	}, []);

    useEffect(() => {
      if (!map) return;

      if (currClickId) {
        console.log("unbind currClickId " + currClickId );
        unByKey(currClickId);
      }

      const newClickId = map.on('singleclick',  onSingleClick);
      setClickId(newClickId);
    }, [mapMode]);

    useEffect(() => {
      if (!map) return;
      console.log("trucks changed! " + JSON.stringify(trucks));

      sourceCar.clear();
      sourceLink.clear();

      trucks.map( t => {
        let tp = transform([t.Long, t.Lat], 'EPSG:4326', 'EPSG:3857');
        let feature = new Feature({
            geometry: new Point( tp ),
            id  : t.Id,
        });
        feature.setId( t.id );
        sourceCar.addFeature(feature);
        if ( t.PointId ) {
          // find wh with this id,
          // take coord and make line
          let w = points.find((el) => el.Id == t.PointId);
          let wp = transform([w.Long, w.Lat], 'EPSG:4326', 'EPSG:3857');
          console.log("t.PointId: " + t.PointId + " " + JSON.stringify(w) + " " + JSON.stringify(wp) );
          let linefeature = new Feature({
            geometry: new LineString([tp, wp]),
            name: "My_Simple_LineString"
          });
          sourceLink.addFeature(linefeature);
        }
      });



    }, [trucks]);

    useEffect(() => {
      if (!map) return;
      console.log("points changed! " + JSON.stringify(points));

      sourceWh.clear();

      points.map( p => {
        let feature = new Feature({
            // long, lat according to specification
            geometry: new Point( transform([p.Long, p.Lat], 'EPSG:4326', 'EPSG:3857')),
            id  : p.Id,
        });
        feature.setId( p.id );
        sourceWh.addFeature(feature);
      });
    }, [points]);

    const onSingleClick = ( evt ) => {
     if (mapMode) {
       const coord = toLonLat(evt.coordinate);
       api.addPoint(coord);
     }
   }

    return (
			<div ref={mapRef} className="ol-map">
				{children}
			</div>
  );
}

//export default Map;
