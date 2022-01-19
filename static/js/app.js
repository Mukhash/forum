(function () {

    const postsEl = document.querySelector('.posts');
    const loaderEl = document.querySelector('.loader');

    // get the quotes from API
    const getQuotes = async (page, limit) => {
        const API_URL = `https://api.javascripttutorial.net/v1/quotes/?page=${page}&limit=${limit}`;
        const response = await fetch(API_URL);
        // handle 404
        if (!response.ok) {
            throw new Error(`An error occurred: ${response.status}`);
        }
        return await response.json();
    }

    // show the quotes
    const showQuotes = (posts) => {
        posts.forEach(post => {
            const postEl = document.createElement('a');
            postEl.href = `post/${post.ID}`
            postEl.classList.add('post');

            postEl.innerHTML = `
            <div class="posts"id=${post.ID}>
                <div class="post-header">
                    <div>
                        <span> <b> ${post.Username} </b></span>
                    </div>
                    <time> ${post.Datefrom} </b></span>
                    </div></time>
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

            postsEl.appendChild(postsEl);
        });
    };

    const hideLoader = () => {
        loaderEl.classList.remove('show');
    };

    const showLoader = () => {
        loaderEl.classList.add('show');
    };

    const hasMoreQuotes = (page, limit, total) => {
        const startIndex = (page - 1) * limit + 1;
        return total === 0 || startIndex < total;
    };

    // load quotes
    const loadQuotes = async (page, limit) => {

        // show the loader
        showLoader();

        // 0.5 second later
        setTimeout(async () => {
            try {
                // if having more quotes to fetch
                if (hasMoreQuotes(page, limit, total)) {
                    // call the API to get quotes
                    const response = await getQuotes(page, limit);
                    // show quotes
                    showQuotes(response.data);
                    // update the total
                    total = response.total;
                }
            } catch (error) {
                console.log(error.message);
            } finally {
                hideLoader();
            }
        }, 500);

    };

    // control variables
    let currentPage = 1;
    const limit = 10;
    let total = 0;


    window.addEventListener('scroll', () => {
        const {
            scrollTop,
            scrollHeight,
            clientHeight
        } = document.documentElement;

        if (scrollTop + clientHeight >= scrollHeight - 5 &&
            hasMoreQuotes(currentPage, limit, total)) {
            currentPage++;
            loadQuotes(currentPage, limit);
        }
    }, {
        passive: true
    });

    // initialize
    loadQuotes(currentPage, limit);

})();