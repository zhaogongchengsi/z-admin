import { defineComponent, h, RendererElement, RendererNode, VNode } from 'vue';
import { ElMenu, ElMenuItem, ElSubMenu, ElMenuItemGroup, ElIcon } from "element-plus"
import type { AsideMenuProps, MenuOptions } from "@/types/index"


type vNode = VNode<RendererNode, RendererElement, {
    [key: string]: any;
}>

export default defineComponent({
    name: 'AsideMenu',
    props: {
        collapse: {
            type: Boolean,
            default: true,
        },
        activeTextColor: {
            type: String,
            default: '#000000',
        },
        backgroundColor: {
            type: String,
            default: '#333333',
        },
        textColor: {
            type: String,
            default: '#ffffff',
        },
        collapseTransition: {
            type: Boolean,
            default: true,
        },
        menuOptions: {
            type: Array,
            default: () => {
                return []
            }
        }
    },
    setup(props, { attrs }) {

        return () => {
            return CreateMenu(props, CreateAsideMenu(props.menuOptions))
        }
    }
})


function CreateAsideMenu(list: MenuOptions[]) :VNode[] {
    return list.map((item: MenuOptions) => {
        // if (!item.disabled) {
        //     return undefined
        // }
        if (item.children && item.children.length > 0) {
            return CreateSubMenu({
                index: item.path,
                path: item.path,
                label: item.label,
            }, CreateAsideMenu(item.children))
        } else {
            return CreateMenuItem({
                path: item.path,
                index: item.path,
            }, item.label)
        }
    })
}

function CreateMenu(props:AsideMenuProps, children: vNode[]) {
    return h(ElMenu, props, {
        default: () => {
            return children
        }
    })
}


interface MenuItemProps {
    index?: any
    label?: string
    path?: string
    disabled?: boolean
}

function CreateMenuItem(props: MenuItemProps, label:any) {
    return h(ElMenuItem, props, {
        default: () => {
           return CreateMenuLabel(label)
        }
    })
}

interface MenuSubProps {
    index: any
    path: string
    label: string
}

function CreateSubMenu(props: MenuSubProps, children: any[]) {
    return h(ElSubMenu, props, {
        default: () => {
            return children.map((item) => CreateMenuItem({
                index: item.path,
                path: item.path,
            },item))
        },
        title: () => {
            return CreateMenuLabel(props.label)
        }
    })
}


function CreateMenuLabel(label: string) {
    return h('span', {}, label)
}