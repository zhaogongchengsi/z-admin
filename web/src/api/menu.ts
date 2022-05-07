import { Get, Post } from "../utils/service";
import { Response } from "@/types/index";


export function getMenuList() {
    return Get(`menu/getmemu`);
}

export function addMenu(data: any) {
    return Post(`menu/createmenu`, data);
}