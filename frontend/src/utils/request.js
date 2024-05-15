import axios from "axios";

const baseURL = "http://localhost:8890";
const instance = axios.create({ baseURL });

instance.interceptors.response.use(
  (result) => {
    return result.data;
  },
  (err) => {
    console.log(err);
    return Promise.reject(err);
  }
);

export default instance;
