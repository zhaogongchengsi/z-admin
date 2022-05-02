import axios from "axios";

// 配置不生效 只能直接想后端发起请求了
const baseURL = "http://localhost:9090/api";

const Http = axios.create({
    baseURL,
    timeout: 5000,
});

export const Get = (url: string) => {
    return new Promise((resolve, reject) => {
        Http.get(url).then(res => {
            resolve(res.data);
        }).catch(err => {
            reject(err);
        });
    })
}

export const Post = (url: string, data: any) => {
    return new Promise((resolve, reject) => {
        Http.post(url, data).then(res => {
            resolve(res.data);
        }).catch(err => {
            reject(err);
        });
    })
}