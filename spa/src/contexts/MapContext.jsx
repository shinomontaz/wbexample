import {createContext, useState, useContext, React} from "react";

// const MapContext = new React.createContext();
// export default MapContext;

const MapContext = createContext({
  map: null,
  setMap: () => {},
  mapMode: false,
  setMode: () => {},
});

export const MapContextProvider = ({ children }) => {
  const [map, setMap] = useState(null);
  const [mapMode, setMode] = useState(false);

  return (
   <MapContext.Provider
     value={{
       map,
       setMap,
       mapMode,
       setMode
     }}
   >
     {children}
   </MapContext.Provider>
 );
};

export const useMapContext = () => useContext(MapContext);
