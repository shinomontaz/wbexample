import axios from "axios";

const axiosClient = axios.create({
  baseURL: `${import.meta.env.VITE_API_BASE_URL}/api`,
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

class Api {
  pause = async () => {
    const res = await axiosClient.post('/pause-fleet');
    return await res.json();
  }

  addPoint = async (location) => {
    const res = await axiosClient.post('/add-point', {location});
    return await res.json();
  }

  generateFleet = async (num, square) => {
    console.log("generateFleet: " + JSON.stringify(square));
    const res = await axiosClient.post('/generate-fleet', {num, square});
    return await res.json();
  }
}

const api = new Api();

export default api;
