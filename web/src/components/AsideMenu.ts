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
        MenuOptions: {
            type: Object,
            default: () => {
                return {}
            }
        }
    },
    setup(props, { attrs }) {

        const menuItemList = props.MenuOptions.map((item: MenuOptions) => {
            if (item.children && item.children.length > 0) {
                return CreateSubMenu(props, item.children)
            } else {
                return CreateMenuItem(props, item)
            }
        })

        return () => {
            return CreateMenu(props, menuItemList)
        }
    }
})

function CreateMenu(props:AsideMenuProps, children: vNode[]) {
    return h(ElMenu, props, children)
}

function CreateMenuItem(props: any, label:any) {
    return h(ElMenuItem, props, { 
        default: () => {
           return CreateMenuLabel(label)
        }
    })
}

function CreateSubMenu(props: any, children: vNode[]) {
    return h(ElSubMenu, props, {
        default: () => {
            return children.map(item => CreateMenuItem(props,item))
        },
        title: (label: string) => {
            return CreateMenuLabel(label)
        }
    })
}


function CreateMenuLabel(label: string) {
    return h('span', {}, label)
}