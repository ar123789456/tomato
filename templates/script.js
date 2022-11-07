class Tomato {
    static Tomato() {

    }

    static registration() {

    }

    static verification() {
        let session = getCookie("tomato-session")
        if (session) {
            this.
        }
    }
}
function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}