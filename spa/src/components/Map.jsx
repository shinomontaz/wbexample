
import React, { useRef, useState, useContext, useEffect } from "react"
import * as ol from "ol";
import TileLayer from "ol/layer/Tile";
import {defaults as defaultControls} from 'ol/control.js';
import OSMSource from "ol/source/OSM";
import { MapContext } from "../contexts/MapContext";


import { transformExtent } from "ol/proj";

const Map = ({ children, zoom, center }) => {
    const mapRef = useRef();

    const { map, setMap } = useContext(MapContext);

    useEffect(() => {
   		let options = {
   			view: new ol.View({ zoom, center }),
   			layers: [
          new TileLayer({
            source: new OSMSource()
          })
        ],
   			controls: [],
   			overlays: []
   		};

   		let mapObject = new ol.Map(options);

      var viewport = transformExtent( mapObject.getView().calculateExtent(mapObject.getSize()), 'EPSG:3857', 'EPSG:4326' );
      console.log("viewport: " + JSON.stringify(viewport));

   		mapObject.setTarget(mapRef.current);
   		setMap(mapObject);

   		return () => mapObject.setTarget(undefined);
   	}, []);

    useEffect(() => {
		    if (!map || Object.keys(map).length == 0 ) return;
        map.getView().setZoom(zoom);



    }, [zoom]);

    useEffect(() => {
    		if (!map || Object.keys(map).length == 0) return;
    		map.getView().setCenter(center)
  	}, [center])

    return (
			<div ref={mapRef} className="ol-map">
				{children}
			</div>
  );
}

export default Map;
