import {createContext, useState, useContext, React} from "react";

// const MapContext = new React.createContext();
// export default MapContext;

const MapContext = createContext({
  map: null,
  viewport: null,
  setMap: () => {},
  vecSource: null,
  setVecSource: () => {},
  mapMode: false,
  setMode: () => {},
  points: [],
  trucks: [],
  positions: null,
  setPoints: () => {},
  setTrucks: () => {},
  setPositions: () => {},
});

export const MapContextProvider = ({ children }) => {
  const [map, setMap] = useState(null);
  const [vecSource, setVecSource] = useState(null);
  const [mapMode, setMode] = useState(false);
  const [viewport, setViewport] = useState(null);

  const [points, setPoints] = useState([]);
  const [trucks, setTrucks] = useState([]);
  const [positions, setPositions] = useState([]);

  return (
   <MapContext.Provider
     value={{
       map,
       setMap,
       vecSource,
       setVecSource,
       viewport,
       setViewport,
       mapMode,
       setMode,
       points,
       trucks,
       positions,
       setPoints,
       setTrucks,
       setPositions,
     }}
   >
     {children}
   </MapContext.Provider>
 );
};

export const useMapContext = () => useContext(MapContext);
