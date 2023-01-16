import store from '../../store';
import Constant from '../../Constant';
import UserService from '../UserService';
import router from '../../router';

export default {

    async authGuardAdmin(to, from, next) {

        try {
            if (localStorage.getItem('token_admin')) {
                const response = await UserService.isAdmin();
                if (response.status==200) {
                    next();
                }
            } else {
                router.push({name:"home"})
            }
        } catch (error) {
            router.push({name:"home"})
            store.dispatch("user/"+Constant.LOGOUT);

        }
    }//authGuardAdmin
}//guards