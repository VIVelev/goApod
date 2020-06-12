<template>
    <div class="root">
        <Loader v-if="!articles" />
        <Sorry v-if="noPost"> Sorry no posts were found!</Sorry>
        <div class="wraper">
            <ArticlePreview
                v-for="article in articles"
                :key="article.id"
                :article="article"
                @click.native="$router.push(`/posts/${article.id}`)"
            />
        </div>
    </div>
</template>

<script>
import Loader from '../UI/Loader'
import Sorry from '../UI/Sorry'
import ArticlePreview from '../Article/ArticlePreview'
export default {
    components: {
        Loader,
        Sorry,
        ArticlePreview,
    },
    data() {
        return {
            articles: null,
            noPost: null,
        }
    },
    async created() {
        const res = await fetch(`${process.env.VUE_APP_API_URL}/articles`)
        let json = await res.json()
        json = [
            {
                id: 1,
                title: 'title today is this mfing title and this is a long one',
                imageUrl:
                    'https://cdn.spacetelescope.org/archives/images/wallpaper2/heic2007a.jpg',
            },
            {
                id: 2,
                title: 'title today is this mfing title',
                imageUrl:
                    'https://images.unsplash.com/photo-1535332371349-a5d229f49cb5?ixlib=rb-1.2.1&w=1000&q=80',
            },
            {
                id: 3,
                title: 'title today is this mfing title',
                imageUrl:
                    'https://cdn.spacetelescope.org/archives/images/wallpaper2/heic2007a.jpg',
            },
            {
                id: 4,
                title: 'title today is this mfing title',
                imageUrl:
                    'https://cdn.spacetelescope.org/archives/images/wallpaper2/heic2007a.jpg',
            },
        ]
        if (!json) {
            setTimeout(() => {
                this.noPost = true
                this.articles = true
            }, 500)
        } else {
            this.articles = json
        }
    },
}
</script>

<style scoped>
.root {
    max-width: 1532px;
    margin: 7rem auto;
}
.wraper {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    grid-template-rows: auto;
    margin: 0.5rem;
}
@media (max-width: 1200px) {
    .wraper {
        grid-template-columns: 1fr 1fr 1fr;
    }
}
@media (max-width: 1000px) {
    .wraper {
        grid-template-columns: 1fr 1fr;
    }
}
@media (max-width: 600px) {
    .wraper {
        grid-template-columns: 1fr;
    }
}
</style>
