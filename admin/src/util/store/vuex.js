import Vue from 'vue'

class Store {
    _vm = new Vue({
        data: {
            $$state: {}
        }, computed: {}
    })

    constructor({ state }) {
        this._vm._data.$$state = state || {}
    }

    set = (name, value) => {
        this.state[name] = value
    }
}

let custom = { state: {} }

custom.state.get = function () {
    // console.log(this._vm)
    return this._vm._data.$$state
}

Object.defineProperties(Store.prototype, custom)

export default {
    Store
}