import { Get, Post } from "../utils/service";
import { Response } from "@/types/index";

export function verify () :Promise<Response<any>> {
    return Get(`other/verify`);
}

export function login <T> (data: T) :Promise<Response<any>> {
    return Post(`user/login`, data);
}