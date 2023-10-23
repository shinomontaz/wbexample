import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import "ol/ol.css";

import { MapContextProvider } from "./contexts/MapContext";

import { Updater } from "./components/Updater";

import router from './router.jsx';
import {RouterProvider} from 'react-router-dom';

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <MapContextProvider>
      <Updater />
      <RouterProvider router={router} />
    </MapContextProvider>
  </React.StrictMode>,
)
