<template>
	<div class="container">

		<div class="row justify-content-center align-content-cener mt-4">

			<div class="col-md-5">
				<div class="container">
					<div class="search">
						<input type="text" class="form-control form-input" v-model="state.filter.search"
							v-on:keyup.enter="search" placeholder="Search for your city...">
						<button type="submit" @click="search">
							<font-awesome-icon class="icon search" icon="fa-solid fa-magnifying-glass" />
						</button>
					</div>
				</div>

			</div>

		</div>

	</div>
</template>

<script>
import { reactive, getCurrentInstance } from 'vue';
import { useRoute } from 'vue-router';

export default {
	setup() {
		const { emit } = getCurrentInstance();
		const currentRoute = useRoute()

		const state = reactive({
			filter: {
				search: ""
			}
		})

		if (currentRoute.params.filter && currentRoute.params.filter != "all") {
			state.filter.search = JSON.parse(atob(currentRoute.params.filter)).search
		}

		const search = () => {
			emit('search', state.filter)
		}
		return { search, state }
	}

}
</script>

<style>
.height {
	height: 100vh;
}

.icon {
	color: white !important;
	width: 25px;
}

.container .search {
	position: relative;
	height: 100%;
}

.container .search input {
	width: 400px;
	height: 50px;
	position: absolute;
	right: 0;
	top: 0;
	padding: 0 25px;
	border: 0;
	border-radius: 25px;
	color: #4e4e4e;
	font-size: 18px;
	font-weight: 300;
	box-shadow: 0 5px 12px 0 rgba(0, 0, 0, 0.1),
		0 1px 28px 0 rgba(0, 0, 0, 0.2),
		0 0 40px 0 rgba(0, 0, 0, 0.1);
	transition: all 0.6s cubic-bezier(0, 2, 1, -1);
}

.container .search button {
	width: 50px;
	height: 50px;
	position: absolute;
	right: 0;
	top: 0;
	margin: auto;
	background: #2d2926;
	color: #fff;
	font-size: 15px;
	border: 0;
	border-radius: 50%;
	box-shadow: 0 6px 28px 0 rgba(0, 0, 0, 0.0), 0 5px 55px 0 rgba(0, 0, 0, 0.0);
	cursor: pointer;
	transition: all 0.6s cubic-bezier(0, 2, 1, -1);
}

@media screen and (max-width : 700px) { 
.container .search {
   position: relative;
   height: 100%;
}
.container .search input {
   width: 50px;
   height: 50px;
   position: absolute;
   right: calc(50% - 25px);
   top: 0;
   padding: 0 25px;
   border: 0;
   border-radius: 25px;
   color: #4e4e4e;
   font-size: 18px;
   font-weight: 300;
   box-shadow: 0 5px 12px 0 rgba(0, 0, 0, 0.1),
       0 1px 28px 0 rgba(0, 0, 0, 0.2),
       0 0 40px 0 rgba(0, 0, 0, 0.1);
   transition: all 0.6s cubic-bezier(0, 2, 1, -1);
}
.container .search button {
   width: 50px;
   height: 50px;
   position: absolute;
   right: calc(50% - 25px);
   top: 0;
   margin: auto;
   background: #2d2926;
   color: #fff;
   font-size: 15px;
   border: 0;
   border-radius: 50%;
   box-shadow: 0 6px 28px 0 rgba(0, 0, 0, 0.0), 0 5px 55px 0 rgba(0, 0, 0, 0.0);
   cursor: pointer;
   transition:all 0.6s cubic-bezier(0, 2, 1, -1);
}
 
.container .search input:focus,
.container .search input:active,
.container .search:hover input {
   width: 400px;
   right:0;
}
.container .search input:focus + button,
.container .search input:active + button,
.container .search:hover button {
   right: 0;
   box-shadow: 0 6px 28px 0 rgba(0, 0, 0, 0.3), 0 5px 55px 0 rgba(0, 0, 0, 0.2);
}

}
</style>
