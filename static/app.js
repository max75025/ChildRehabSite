new Vue({
    el: "#app",
    data: {
        posts: []
    },
    created() {
        fetch("https://jsonplaceholder.typicode.com/posts/8")
            .then(response => {
                if (response.ok) {
                    return response.json();
                }

                throw new Error("Network response was not ok");
            })
            .then(json => {
                this.posts.push({
                    title: json.title,
                    body: json.body
                });
            })
            .catch(error => {
                console.log(error);
                alert(error);
            });
    }
});
