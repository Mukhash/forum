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
                        <span style="color: #518fa1;"> 🗩 </span>
                        <span> ${post.CommentsCount} </span>
                    </div>

                    <form class="post-rating" action="/like" method="POST">
                        <input name="objType" value="1" type="hidden">
                        <input name="objID" value="${post.ID}" type="hidden">

                        <button class="rateButton" type="submit" name="action" value="1" style="background-color: #CDF2CA;"> ⮝
                        </button>

                        <span class="rating"> ${post.LikesCount} </span>

                        <button class="rateButton" type="submit" name="action" value="2"style="background-color: #FFDEFA;"> ⮟
                        </button>
                    </form>
                </div>
            </div>
            `;

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
        if (id === 0)
        return false
        else
        return true
    };

    const loadQuotes = async (limit) => {

        //showLoader();

        setTimeout(async () => {
            try {
                // if having more quotes to fetch
                if (hasMoreQuotes(firstID)) {

                    const response = await getQuotes(firstID, limit);

                    showQuotes(response.data);

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
            hasMoreQuotes(firstID)) {
            loadQuotes(firstID, limit);
        }
    }, {
        passive: true
    });

    // initialize
    loadQuotes(firstID, limit);

})();