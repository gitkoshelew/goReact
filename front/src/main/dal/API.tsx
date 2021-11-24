import axios from 'axios';

export class PetHotelsService {
    _apiBase = 'https://reqres.in/';


    getAllHotels = async () => {
        const res = await axios.get(this._apiBase + 'api/users?page=1')
       return this._transform(res)
    }


    _transform = (res: any) => {
        return res.data.data;
    }

}