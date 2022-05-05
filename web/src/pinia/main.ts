import { defineStore } from 'pinia'

export const mainStore = defineStore("main", {
    state: () => {
        return {
            collapse: true,
            theme: {
                backgroundColor: '#ffffff',
                textColor: '#333333',
                activeTextColor: '#000000',
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