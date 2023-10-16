
import React, { useRef, useState, useContext, useEffect } from "react"
import * as ol from "ol";
import TileLayer from "ol/layer/Tile";
import VectorLayer from 'ol/layer/Vector';
import VectorSource from 'ol/source/Vector';

import {defaults as defaultControls} from 'ol/control.js';
import OSMSource from "ol/source/OSM";
import { useMapContext } from "../contexts/MapContext";

import {unByKey} from "ol/Observable";

import { toLonLat, transformExtent } from "ol/proj";
import {toStringHDMS} from 'ol/coordinate.js';

import api from '../api';

export default function Map({ children, zoom, center, mapMode }) {
    const mapRef = useRef();

    const { map, setMap } = useMapContext();
    const [currClickId, setClickId ] = useState(null);

    useEffect(() => {
   		let options = {
   			view: new ol.View({ zoom, center }),
   			layers: [
          new TileLayer({
            source: new OSMSource()
          }),
          new VectorLayer({
            source: new VectorSource()
          }),
        ],
   			controls: [],
   			overlays: []
   		};

   		let mapObject = new ol.Map(options);

      var viewport = transformExtent( mapObject.getView().calculateExtent(mapObject.getSize()), 'EPSG:3857', 'EPSG:4326' );

   		mapObject.setTarget(mapRef.current);

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
