import { Get, Post } from "../utils/service";
import { Response } from "@/types/index";


export function getMenuList() :Response {
    return Get(`menu/getmemu`);
}

export function addMenu(data: any) :Response {
    return Post(`menu/createmenu`, data);
}
