import { defineStore } from "pinia";
import { getMenuList } from "@/api/menu";
import { routerToTree } from "@/utils/index";

import type { RouterInfo } from "@/types/index";

const components = import.meta.glob("../views/**/*.vue");



// 优化路由组件列表
function getRouters() {
  const routerMap = new Map();
  for (let key in components) {
    let _key = key.replace(/^\..\//, "");
    routerMap.set(_key, components[key]);
  }
  return routerMap;
}

function getAsyncComponents(routers: any[], components: any) {
  return routers.map((item) => {
    if (item.component) {
      // 获取组件
      let _component = components.get(item.component);
      if (_component) {
        // 组件存在 则 替换
        item.component = _component;
      } else {
        // 不存在则替换成 报错组件 或者直接删除掉当前路由
        // item.component = () => import("../views/NotFon.vue");
      }
    }
    if (item.children) {
      // 当子路由存在时，递归处理
      item.children = getAsyncComponents(item.children, components);
    }
    return item;
  });
}


export const routerStore = defineStore("router", {
  state: () => {
    return {
      // router 原始数据
      routerInfo: [],
    };
  },

  getters: {
    getRouteData () :RouterInfo[] {
        return routerToTree(this.routerInfo);
    },
    getRouter ( ) :any[] {
        const componentsModels = getRouters()

        const r = JSON.parse(JSON.stringify(this.routerInfo));

        return getAsyncComponents(routerToTree(r), componentsModels);
    },
  },

  actions: {
    async getRouterInfo() {
      const routerInfo = await getMenuList();

      if (routerInfo.code === 200) {
        this.routerInfo = routerInfo.data;
        return true;
      } else {
        return false;
      }
    },
  },
});
