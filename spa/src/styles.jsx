import { Circle, Fill, Stroke, Style } from "ol/style";

class Styles {
  cache = {};

  Get = (kind, type) => {

  }

  getIcon = ( icon, size, color, border ) => {
    if( !this.cache['getIcon' + icon + size + color + border] ) {
        this.cache['getIcon' + icon + size + color + border] = new Style({
            image: new Circle({
                radius: size - 1,
                fill: new Fill({
                    color: color,
                }),
                stroke: new Stroke({
                    width: 1,
                    color: border
                }),
            }),
            text: new ol.style.Text({
                text: icon,
                font : 'Normal '+(size)+'px FontAwesome',
                fill: new ol.style.Fill({
                    color: '#fff',
                })
            })
        });
    }
    return this.cache['getIcon' + icon + size + color + border];
  }
}
const styles = new Styles();

export default styles;
