import { defineStore } from 'pinia'

export const mainStore = defineStore("main", {
    state: () => {
        return {
            collapse: true,
            theme: {
                backgroundColor: '#333333',
                textColor: '#ffffff',
                activeTextColor: '#cccccc',
            },
            loading: false,
        }
    },

    getters: {
        backgroundColor: state => state.theme.backgroundColor,
        textColor: state => state.theme.textColor,
        activeTextColor: state => state.theme.activeTextColor,
        coll: state => state.collapse,
        loading: state => state.loading,
    },

    actions: {
        toggleCollapse () {
            this.collapse = !this.collapse;
        }
    },
})