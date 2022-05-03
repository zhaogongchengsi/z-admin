import { Get, Post } from "../utils/service";


export const verify = () => {
    return Get(`other/verify`);
}

export const login = (data: any) => {
    return Post(`user/login`, data);
}