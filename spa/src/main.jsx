import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import "ol/ol.css";

import "./assets/fontawesome/css/all.min.css";

import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'

library.add(fas)

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
