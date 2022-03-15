(function () {

    const postsEl = document.querySelector('.posts_feed');
   // const loaderEl = document.querySelector('.loader');



    const getQuotes = async (first_id, limit) => {
        const API_URL = `http://localhost:8080/next_posts?first_id=${first_id}&limit=${limit}`;
        const response = await fetch(API_URL);
        // handle 404
        if (!response.ok) {
            throw new Error(`An error occurred: ${response.status}`);
        }
        return await response.json();
    }

    const showQuotes = (posts) => {
        posts.forEach(post => {
            const postEl = document.createElement('a');
            postEl.href = `post/${post.ID}`

            postEl.innerHTML = `
            <div class="post" id=${post.ID}>
                <div class="post-header">
                    <div>
                        <span> <b> ${post.Username} </b></span>
                    </div>
                    <time> ${post.Datefrom} </b> </time>
                </div>

                <pre>${post.Body}</pre>

                <div class="post-bottom">

                    <div>
                        <span style="color: #518fa1;"><i class="fa-solid fa-comments"></i></span>
                        <span> ${post.CommentsCount} </span>
                    </div>

                    <form class="post-rating" action="/like_post" method="post">
                        <input name="post_id" value="${post.ID}" type="hidden">
                        <button class="likeButton" type="submit" name="action" value="1"><i class="fa-solid fa-thumbs-up"></i>
                        </button>

                        <span class="rating"> ${post.LikesCount} </span>

                        <button class="dislikeButton" type="submit" name="action" value="2"><i class="fa-solid fa-thumbs-down"></i>
                        </button>
                    </form>
                </div>
            </div>
            `;
            postEl.st

            postsEl.appendChild(postEl);
        });
    };

    // const hideLoader = () => {
    //     loaderEl.classList.remove('show');
    // };

    // const showLoader = () => {
    //     loaderEl.classList.add('show');
    // };

    const hasMoreQuotes = (id) => {
        return id !== 0;
    };

    const loadQuotes = async (first_id, limit) => {

        //showLoader();

        setTimeout(async () => {
            try {
                // if having more quotes to fetch
                if (hasMoreQuotes(first_id)) {

                    const response = await getQuotes(first_id, limit);

                    showQuotes(response.data);

                    firstID = response.nextFirstId;
                    console.log(response.nextFirstId);
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
            hasMoreQuotes(firstID)) {
            loadQuotes(firstID, limit);
        }
    }, {
        passive: true
    });

    // initialize
    loadQuotes(firstID, limit);

})();