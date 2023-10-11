import {createContext, useState, React} from "react";

// const MapContext = new React.createContext();
// export default MapContext;

export const MapContext = createContext({
  map: {},
  setMap: () => {},
});

export const MapContextProvider = ({ children }) => {
  const [map, setMap] = useState({});

  return (
   <MapContext.Provider
     value={{
       map,
       setMap
     }}
   >
     {children}
   </MapContext.Provider>
 );
};
