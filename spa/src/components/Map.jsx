
import React, { useRef, useState, useContext, useEffect } from "react"
import * as ol from "ol";
import TileLayer from "ol/layer/Tile";
import VectorLayer from 'ol/layer/Vector';
import VectorSource from 'ol/source/Vector';
import Feature from 'ol/Feature';
import { Circle as CircleStyle, Fill, Stroke, Style } from "ol/style";

import {Point, Polygon} from 'ol/geom';
import {defaults as defaultControls} from 'ol/control.js';
import OSMSource from "ol/source/OSM";
import { useMapContext } from "../contexts/MapContext";

import {unByKey} from "ol/Observable";

import { transform, fromLonLat, toLonLat, transformExtent } from "ol/proj";
import {toStringHDMS} from 'ol/coordinate.js';

import api from '../api';

export default function Map({ children, zoom, center, mapMode }) {
    const mapRef = useRef();

    const { map, setMap, vecSource, setVecSource, setViewport, points, trucks, positions } = useMapContext();
    const [currClickId, setClickId ] = useState(null);

    useEffect(() => {
      let sVec = new VectorSource();
   		let options = {
   			view: new ol.View({ zoom, center }),
   			layers: [
          new TileLayer({
            source: new OSMSource()
          }),
          new VectorLayer({
            source: sVec
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

      setVecSource(sVec);
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
      console.log("positions changed! " + JSON.stringify(positions))
    }, [positions]);

    useEffect(() => {
      if (!map) return;
      console.log("trucks changed! " + JSON.stringify(trucks))
    }, [trucks]);

    useEffect(() => {
      if (!map) return;
      console.log("points changed! " + JSON.stringify(points));

      points.map( p => {
        let feature = new Feature({
            // long, lat according to specification
            geometry: new Point( transform([p.Long,  p.Lat], 'EPSG:4326', 'EPSG:3857')),
            id  : p.Id,
            style: new Style({
    image: new CircleStyle({
      radius: 10,
      fill: null,
      stroke: new Stroke({
        color: "magenta",
      }),
    }),
  }),
        });
        feature.setId( p.id );
        vecSource.addFeature(feature);
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
