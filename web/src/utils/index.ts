



// 排序
export function sortByRouter(router: any[]) {
  return router.sort((a, b) => {
    return a.ID - b.ID;
  });
}

// 路由数据转变成树形结构
export function routerToTree(router: any[]) {
  // 筛选出顶级菜单 并且排序
  let top_menus = sortByRouter(
    router.filter((item) => {
      return item.parent_id === 0;
    })
  );
  // 筛选出顶级菜单下的子菜单
  let children_menus = router.filter((item) => {
    return item.parent_id !== 0;
  });

  // 将子菜单挂载到顶级菜单下
  function subscribeToChildren(top_menus: any[], children_menus: any[]) {
    top_menus.forEach((item) => {
      // 每次都要排序一下先
      let _child = sortByRouter(
        children_menus.filter((child) => {
          return child.parent_id === item.ID;
        })
      );

      if (_child.length > 0) {
        item.children = _child;
        subscribeToChildren(item.children, children_menus);
      }
    });
  }

  subscribeToChildren(top_menus, children_menus);
  return top_menus;
}


