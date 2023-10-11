
import React, { useRef, useState, useEffect, useContext } from "react"
import { fromLonLat, toLonLat, get } from "ol/proj";
import Map from '../components/Map'
import { MapContextProvider, MapContext } from "../contexts/MapContext";

import {PauseCircleIcon, PlusCircleIcon} from '@heroicons/react/24/outline';

import api from '../api';

const center = [37.5, 55.7];
const zoom = 10;

const Landmap = () => {

  const { map } = useContext(MapContext);
  var viewport;

  useEffect(() => {
      if (!map || Object.keys(map).length == 0 ) return;

  });

  const onGenerateFleet = () => {
    api.generateFleet(10, viewport);
  }

    return (
          <>
          <header className="bg-white shadow">
            <div className="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
            <h1 className="text-3xl font-bold tracking-tight text-gray-900">Map</h1>
            </div>
          </header>
          <main>
            <div className="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
              <div className="mt-2 flex gap-x-6">
                <a
                  href="#"
                  className="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                  onClick={()=> onGenerateFleet()}
                >
                  Generate trucks
                </a>
                <button
                  type="button"
                  className="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700"
                >
                  <span className="sr-only">Open main menu</span>
                  <PauseCircleIcon className="h-6 w-6" aria-hidden="true" />
                </button>
                <button
                  type="button"
                  className="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700"
                >
                  <span className="sr-only">Add point</span>
                  <PlusCircleIcon className="h-6 w-6" aria-hidden="true" />
                </button>
              </div>
              <MapContextProvider>
              <Map center={fromLonLat(center)} zoom={zoom} />
              </MapContextProvider>
            </div>
          </main>
          </>
  );
}

export default Landmap;