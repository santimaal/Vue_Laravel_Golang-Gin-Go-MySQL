import Vuex from "vuex";
// import Constant from '../Constant';
// import shortid from 'shortid';
import { table } from './modules/table'
import { thematic } from "./modules/thematic";

export default Vuex.createStore({
    modules : {
        table : table,
        thematic: thematic
    }
});