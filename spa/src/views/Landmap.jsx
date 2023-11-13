
import React, { useRef, useState, useEffect, useContext } from "react"
import { fromLonLat, toLonLat, get } from "ol/proj";
import Map from '../components/Map'
import { MapContextProvider, useMapContext } from "../contexts/MapContext";

import {PlayCircleIcon, PauseCircleIcon, PlusCircleIcon, ArrowPathIcon, ArrowPathRoundedSquareIcon, ExclamationTriangleIcon} from '@heroicons/react/24/outline';
import {PlusCircleIcon as PlusCircleIconSolid} from '@heroicons/react/24/solid';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

import api from '../api';

const center = [37.5, 55.7];
const zoom = 10;

const Landmap = () => {
    const { map, viewport, mapMode, setMode, setPoints, setTrucks } = useMapContext();

    const onStartFleet = () => {
      api.start();
    }
    const onPauseFleet = () => {
      api.pause();
    }

    const onUpdatePoints = async () => {
      await api.getPoints( ).then( (points) => {
          console.log( "Points received: " + JSON.stringify(points) );
          setPoints(points);
        }
      ).catch( (err) => {
        console.log( "getPoints exception " + JSON.stringify(err) );
        }
      );
    }
    const onUpdateFleet = async () => {
      await api.getTrucks( ).then( (trucks) => {
          console.log( "Trucks received: " + JSON.stringify(trucks) );
          setTrucks(trucks);
        }
      ).catch( (err) => {
        console.log( "getTrucks exception " + JSON.stringify(err) );
        }
      );
    }

    const onGenerateFleet = () => {
      api.generateFleet(10, viewport);
    }

    const onPointAddMode = () => {
      if (!map || Object.keys(map).length == 0) return;
      setMode(!mapMode);
    }

    const onSolve = () => {
      api.solve();
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
                  onClick={()=> onStartFleet()}
                >
                  <span className="sr-only">Start fleet</span>
                  <PlayCircleIcon className="h-6 w-6" aria-hidden="true" />
                </button>
                <button
                  type="button"
                  className="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700"
                  onClick={()=> onPauseFleet()}
                >
                  <span className="sr-only">Pause fleet</span>
                  <PauseCircleIcon className="h-6 w-6" aria-hidden="true" />
                </button>
                <button
                  type="button"
                  className="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700"
                  onClick={()=> onPointAddMode()}
                >
                  <span className="sr-only">Add point</span>
                  {mapMode ? <PlusCircleIconSolid className="h-6 w-6" aria-hidden="true" /> : <PlusCircleIcon className="h-6 w-6" aria-hidden="true" /> }
                </button>
                <button
                  type="button"
                  className="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700"
                  onClick={()=> onUpdatePoints()}
                >
                  <span className="sr-only">Update points</span>
                  <ArrowPathIcon className="h-6 w-6" aria-hidden="true" />
                </button>
                <button
                  type="button"
                  className="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700"
                  onClick={()=> onUpdateFleet()}
                >
                  <span className="sr-only">Update fleet</span>
                  <ArrowPathRoundedSquareIcon className="h-6 w-6" aria-hidden="true" />
                </button>

                <button
                  type="button"
                  className="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700"
                  onClick={()=> onSolve()}
                >
                  <span className="sr-only">Solve!</span>
                  <ExclamationTriangleIcon className="h-6 w-6" aria-hidden="true" />
                </button>
              </div>

              <Map center={fromLonLat(center)} zoom={zoom} mapMode={mapMode}/>
            </div>
          </main>
          </>
  );
}

export default Landmap;
