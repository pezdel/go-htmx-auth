import Alpine from 'alpinejs'


window.app = function() {
    return {
        name: "Adam",
        age: 34,
        sidebar: false,
        toggle(){
            this.sidebar = !this.sidebar
        },
    }
}

Alpine.start();
