<template>
    <article class="card one_card">
        <div class="card__info-hover">
            <button>Reserve</button>
            <div class="price">
                <p>29'99€</p>
            </div>
        </div>
        <div class="card__img"></div>
        <a href="#" class="card_link">
            <div class="card__img--hover"></div>
        </a>
        <div class="card__info">
            <h3 class="card__title">Table Sanvic {{ tableitem.id }}</h3>
            <div class="container">
                <div class="row">
                    <div v-for="(item, id) in state.thematic" :key="id">
                        <p v-if="(item.id == tableitem.id_thematic)"><font-awesome-icon class="rIcon"
                                icon="fa-solid fa-utensils" /> {{ item.name }}</p>
                    </div>
                    <p class="cpacity"><font-awesome-icon class="rIcon" icon="fa-solid fa-person" />{{
                            tableitem.capacity
                    }}</p>
                </div>
                <div class="row">
                    <div v-for="(item, id) in state.thematic" :key="id">
                        <p v-if="(item.id == tableitem.id_thematic)"><font-awesome-icon class="rIcon"
                                icon="fa-solid fa-location-dot" />{{ item.location }}</p>
                    </div>
                    <div class="loc_rest">
                        <p v-if="tableitem.location === 'outside'"><font-awesome-icon class="rIcon"
                                icon="fa-solid fa-person-shelter" />Outside</p>
                        <p v-if="tableitem.location === 'inside'"><font-awesome-icon class="rIcon"
                                icon="fa-solid fa-person-shelter" /> Inside</p>
                    </div>
                </div>
            </div>


        </div>
    </article>
</template>

<script>
import { useStore } from 'vuex';
import { reactive, computed } from 'vue';

export default {
    props: {
        tableitem: Object
    },
    setup() {
        const store = useStore();


        const state = reactive({
            thematic: computed(() => store.getters["thematic/getThematic"])
        })

        return { state }
    }
}
</script>

<style lang="scss" scoped>
@import url("https://fonts.googleapis.com/css2?family=Lobster&display=swap");


.rIcon {
    height: 25px;
    color: black !important;
    margin-right: 15px;
}

.one_card .card__img,
.one_card .card__img--hover {
    background-image: url('https://architecturesideas.com/wp-content/uploads/2020/10/Guide-to-Restaurant-Tables-5.jpg');
}

.price {
    float: right;
}

.card__img {
    visibility: hidden;
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    width: 100%;
    height: 235px;
    border-top-left-radius: 12px;
    border-top-right-radius: 12px;
}

.card__info-hover {
    position: absolute;
    padding: 16px;
    width: 100%;
    opacity: 0;
    top: 0;
}

.card__img--hover {
    transition: 0.2s all ease-out;
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    width: 100%;
    position: absolute;
    height: 235px;
    border-top-left-radius: 12px;
    border-top-right-radius: 12px;
    top: 0;
}

.card {
    transition: all .4s cubic-bezier(0.175, 0.885, 0, 1);
    background-color: #fff;
    width: 100%;
    position: relative;
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0px 13px 10px -7px rgba(0, 0, 0, 0.1);

    &:hover {
        box-shadow: 0px 30px 18px -8px rgba(0, 0, 0, 0.1);
        transform: scale(1.10, 1.10);

        .card__img--hover {
            height: 100%;
            opacity: 0.3;
        }

        .card__info {
            background-color: transparent;
            position: relative;
        }

        .card__info-hover {
            opacity: 1;
        }
    }
}

.card__info {
    z-index: 2;
    background-color: #fff;
    border-bottom-left-radius: 12px;
    border-bottom-right-radius: 12px;
    padding: 16px 24px 24px 24px;
    font-size: large;
    .cpacity {
    margin-left: 8%;
    }
    .loc_rest{
    margin-left: 35%;
    }
}

.card__title {
    margin-top: 5px;
    margin-left: -10%;
    margin-bottom: 20px;
    text-decoration: underline;
    font-style: italic;
    text-align: center;
    font-family: "Lobster", cursive;
}
</style>