<?php

namespace App\Http\Requests\User;

use Illuminate\Foundation\Http\FormRequest;

class StoreUserRequest extends FormRequest
{
    // Authorize (return boolean)    
    public function authorize()
    {
        return true;
    }

    // Rules (return array)
    public function rules()
    {
        return [
            "is_active" => ["optional"],
            "name" => ["required"],
            "password" => ["required"],
            "type" => ["optional"],
            "email" => ["required"],
            "img" => ["optional"],
            "noti" => ["optional"],
            "chat_id" => ["optional"],
        ];
    }
}
