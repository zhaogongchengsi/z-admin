import { FormItemRule } from "element-plus";


export interface Login {
    username: string;
    password: string;
    verify: string;
    verifyId?: string
    [key: string]: any
}

export interface verify {
    verifyId: string;
    verify: string;
}

export interface LoginRes {
    user: any;
    token: string;
}

export const rules = {
    username: {
        type: 'string',
        required: true,
        message: '请输入用户名',
    },
    password: {
        type: 'string',
        required: true,
        message: '请输入密码',
    },
    verify: {
        type: 'string',
        required: true,
        message: '请输入验证码',
    }
}

