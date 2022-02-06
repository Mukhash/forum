(function () {

    const commentsEl = document.querySelector('.comments');
    const post_id = commentsEl.id
   // const loaderEl = document.querySelector('.loader');

    const getData = async (post_id, first_id, limit) => {
        const API_URL = `http://localhost:8080/next_comments?post_id=${post_id}&first_id=${first_id}&limit=${limit}`;
        const response = await fetch(API_URL);
        // handle 404
        if (!response.ok) {
            throw new Error(`An error occurred: ${response.status}`);
        }
        return await response.json();
    }

    const showData = (comments) => {
        comments.forEach(comment => {
            const commentEl = document.createElement('div');
            commentEl.classList.add("comment")
            commentEl.id = `${comment.ID}`

            commentEl.innerHTML = `
            <a href="${comment.ID}"></a>
            <div class="comment-header">
                <span> <b>${comment.Username} </b> </span>
                <time> ${comment.Datefrom}</time>
            </div>

            <pre>${comment.Body}</pre>

            <form class="comment-rating" action="/like" method="POST">
                <input name="objType" value="2" type="hidden">
                <input name="objID" value="" type="hidden">

                <button class="rateButton" type="submit" name="action" value="1" 
                    style="background-color: #CDF2CA;"> ⮝ </button>

                <span class="rating"> </span>

                <button class="rateButton" type="submit" name="action" value="2"
                    style="background-color: #FFDEFA;"> ⮟ </button>
            </form>
            `;

            commentsEl.appendChild(commentEl);
        });
    };

    // const hideLoader = () => {
    //     loaderEl.classList.remove('show');
    // };

    // const showLoader = () => {
    //     loaderEl.classList.add('show');
    // };

    const hasMoreData = (id) => {
        if (id === 0)
        return false
        else
        return true
    };

    const loadData = async (limit) => {

        //showLoader();

        setTimeout(async () => {
            try {
                // if having more quotes to fetch
                if (hasMoreData(firstID)) {

                    const response = await getData(post_id, firstID, limit);

                    showData(response.data);

                    firstID = response.nextFirstId;
                }
            } catch (error) {
                console.log(error.message);
            } finally {
                // hideLoader();
            }
        }, 500);

    };

    // control variables
    let firstID = -1
    const limit = 15;

    window.addEventListener('scroll', () => {
        const {
            scrollTop,
            scrollHeight,
            clientHeight
        } = document.documentElement;

        if (scrollTop + clientHeight >= scrollHeight - 5 &&
            hasMoreData(firstID)) {
            loadData(firstID, limit);
        }
    }, {
        passive: true
    });

    // initialize
    loadData(firstID, limit);

})();