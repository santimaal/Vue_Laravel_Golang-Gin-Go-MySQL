<?php

namespace App\Http\Requests\Reserve;

use Illuminate\Foundation\Http\FormRequest;

class StoreReserveRequest extends FormRequest
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
            "id_table" => ["optional"],
            "id_user" => ["optional"],
            "is_confirmed" => ["required"],
            "dateini" => ["optional"],
        ];
    }
}