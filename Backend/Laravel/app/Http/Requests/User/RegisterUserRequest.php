<?php

namespace App\Http\Requests\User;

use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Http\Exceptions\HttpResponseException;


class RegisterUserRequest extends FormRequest
{
    // Authorize (return boolean)    
    public function authorize()
    {
        return true;
    }

    // Rules (return array)
    public function rules(){
        return [
            "is_active" => 'optional',
            "name" => 'required|string|between:2,100|unique:users',
            "password" => 'required|string|min:2',
            "type" => 'optional',
            "email" => 'required|string|email|max:100|unique:users',
            "img" => 'optional',
            "chat_id" => 'optional',
        ];
    }

    public function withValidator($validator) {
        if ($validator->fails()) {
            throw new HttpResponseException(response()->json([
                'errors' => $validator->errors(),
            ], 400));
        } else {
            return "succesfully";
        }
    }
       
}
