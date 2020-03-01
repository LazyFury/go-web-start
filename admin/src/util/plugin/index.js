import Vue from 'vue';
import clipBoard from './clipBroad/clipBoard'
import aplit_array from './func/aplit_array'
import money_format from './func/money_format'
import time_format from './func/time_format'
import phoneCall from './phoneCall'

let mixins = {

}

let func = {
    h5: {
        clipBoard,
        phoneCall
    },
    every: {
        arrayAplit: aplit_array,
        moneyFormat: money_format,
        timeFormat: time_format
    },
    uni: {

    }
}

let config = {

}

function install(Vue) {
    Vue.prototype.$custom = {
        name: "custom_plugin",
        version: "1.0.0",
        mixins,
        func,
        config
    }
}

export default {
    install
}