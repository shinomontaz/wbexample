import axios from "axios";

const axiosClient = axios.create({
  baseURL: `${import.meta.env.VITE_API_BASE_URL}/spa`,
});

// axiosClient.interceptors.request.use((config) => {
//   config.headers.Authorization = `Bearer ${localStorage.getItem('TOKEN')}`
//   return config
// });
//
// axiosClient.interceptors.response.use(response => {
//   return response;
// }, error => {
//   if (error.response && error.response.status === 401) {
//     localStorage.removeItem('TOKEN')
//     window.location.reload();
//     return error;
//   }
//   throw error;
// })


//http://wb.loc:8080/spa/PauseFleet
class Api {
  getPoints = async () => {
    const res = await axiosClient.get('/points');
    console.log("res: " + JSON.stringify(res));
    return res.data.data;
  }
  getTrucks = async () => {
    const res = await axiosClient.get('/trucks');
    return await res.data;
  }
  getPositions = async () => {
    const res = await axiosClient.get('/positions');
    return await res.data;
  }

  pause = async () => {
    const res = await axiosClient.post('/pause-fleet');
    return await res.data;
  }

  addPoint = async (location) => {
//    console.log(KSON.stringify(location))
    const res = await axiosClient.post('/add-point', {location});
    return await res.data;
  }

  generateFleet = async (num, square) => {
    console.log("generateFleet: " + JSON.stringify(square));
    const res = await axiosClient.post('/generate-fleet', {num, square});
    return await res.data;
  }
}

const api = new Api();

export default api;
