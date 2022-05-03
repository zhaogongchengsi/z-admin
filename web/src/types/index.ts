import { AxiosResponse } from "axios";



export interface Response <T> {
    code: number;
    data: T;
    msg: string;
}

