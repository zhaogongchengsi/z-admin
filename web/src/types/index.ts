
export interface Response <T> {
    code: number;
    data: T;
    msg: string;
}


export interface MenuOptions {
    label: string
    icon?: string
    path?: string
    children?: MenuOptions[]
}

export interface AsideMenuProps {
    collapse: boolean
    activeTextColor: string
    textColor: string
    backgroundColor: string
    collapseTransition: boolean
    MenuOptions: MenuOptions[]
}