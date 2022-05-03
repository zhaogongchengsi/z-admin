import axios from "axios";
import { Response } from "../types/index";

// 配置不生效 只能直接想后端发起请求了
const baseURL = "http://localhost:9090/api";

const Http = axios.create({
    baseURL,
    timeout: 5000,
});


// 响应拦截器
Http.interceptors.response.use(
    (response: any) => {
        return response;
    },
    (error: any) => {
        return Promise.reject(error);
    }
)


// 请求拦截器
Http.interceptors.request.use(
    (config) => {
        config.headers['token'] = `${localStorage.getItem("z_token")}`;
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
)

export const Get = (url: string) => {
    return new Promise<Response<any>>((resolve, reject) => {
        Http.get(url).then(res => {
            resolve(res.data);
        }).catch(err => {
            reject(err) ;
        });
    })
}

export const Post = (url: string, data: any) => {
    return new Promise<Response<any>>((resolve, reject)  => {
        Http.post(url, data).then(res => {
            resolve(res.data);
        }).catch(err => {
            reject(err);
        });
    })
}