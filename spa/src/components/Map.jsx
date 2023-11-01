
import React, { useRef, useState, useContext, useEffect } from "react"
import * as ol from "ol";
import TileLayer from "ol/layer/Tile";
import VectorLayer from 'ol/layer/Vector';
import VectorSource from 'ol/source/Vector';
import Feature from 'ol/Feature';
import { Circle, Fill, Stroke, Text, Style } from "ol/style";

import {Point, Polygon} from 'ol/geom';
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

    const { map, setMap, sourceWh, setWhSource, sourceCar, setCarSource, setViewport, points, trucks, positions } = useMapContext();
    const [currClickId, setClickId ] = useState(null);

    useEffect(() => {
      let sWhVec = new VectorSource();
      let sCarVec = new VectorSource();
   		let options = {
   			view: new ol.View({ zoom, center }),
   			layers: [
          new TileLayer({
            source: new OSMSource()
          }),
          new VectorLayer({
            source: sWhVec,
            style: (feature) => {
              return new Style({
                image: new Circle({
                radius: 10,
                fill: new Fill({
                    color: '#d43f3a',
                }),
                stroke: new Stroke({
                    width: 1,
                    color: '#000'
                }),
                }),
                text: new Text({
                  text: '\uf468',
                  font: 'normal 18px FontAwesome',
                  textBaseline: 'Bottom',
                  fill: new Fill({
                    color: 'white',
                  })
                })
              });
            }
          }),
          new VectorLayer({
            source: sCarVec,
            style: (feature) => {
              return new Style({
                image: new Circle({
                radius: 10,
                fill: new Fill({
                    color: '#d43f3a',
                }),
                stroke: new Stroke({
                    width: 1,
                    color: '#000'
                }),
                }),
                text: new Text({
                  text: '\uf48b',
                  font: 'normal 16px FontAwesome',
                  textBaseline: 'Bottom',
                  fill: new Fill({
                    color: 'white',
                  })
                })
              });
            }
          }),
        ],
   			controls: [],
   			overlays: []
   		};

   		let mapObject = new ol.Map(options);

      var viewport = transformExtent( mapObject.getView().calculateExtent(mapObject.getSize()), 'EPSG:3857', 'EPSG:4326' );
      setViewport(viewport);
      console.log(viewport);

   		mapObject.setTarget(mapRef.current);

      setWhSource(sWhVec);
      setCarSource(sCarVec);

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
      console.log("positions changed! " + JSON.stringify(positions));

      points.map( p => {
        let feature = new Feature({
            // long, lat according to specification
            geometry: new Point( transform([p.Long,  p.Lat], 'EPSG:4326', 'EPSG:3857')),
            type: 1, // TODO: type of a WH
            id  : p.Id,
        });
        feature.setId( p.id );
        sourceCar.addFeature(feature);
      });

    }, [positions]);

    useEffect(() => {
      if (!map) return;
      console.log("trucks changed! " + JSON.stringify(trucks));

      // here we need add info of coupling car to wh

    }, [trucks]);

    useEffect(() => {
      if (!map) return;
      console.log("points changed! " + JSON.stringify(points));

      points.map( p => {
        let feature = new Feature({
            // long, lat according to specification
            geometry: new Point( transform([p.Long,  p.Lat], 'EPSG:4326', 'EPSG:3857')),
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
