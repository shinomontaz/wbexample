import { useContext, useEffect } from "react";
import MapContext from "../contexts/MapContext";

import {Control, defaults as defaultControls} from 'ol/control.js';
import { FullScreen } from "ol/control";

// const AddPOIControl = () => {
//   const { map } = useContext(MapContext);
//
//   useEffect(() => {
//   		if (!map) return;
//
//   		let fullScreenControl = new FullScreen({});
//
//   		map.controls.push(fullScreenControl);
//
//   		return () => map.controls.remove(fullScreenControl);
//   	}, [map]);
//
//   	return null;
// }



class AddPOIControl extends Control {
  /**
   * @param {Object} [opt_options] Control options.
   */
  constructor(opt_options) {
    const { map } = useContext(MapContext);
    const options = opt_options || {};

    const button = document.createElement('button');
    button.innerHTML = 'N';

    const element = document.createElement('div');
    element.className = 'rotate-north ol-unselectable ol-control';
    element.appendChild(button);

    super({
      element: element,
      target: options.target,
    });

    button.addEventListener('click', this.handleRotateNorth.bind(this), false);
  }

  handleRotateNorth() {
    this.map.getView().setRotation(0);
  }
}

export default AddPOIControl;
