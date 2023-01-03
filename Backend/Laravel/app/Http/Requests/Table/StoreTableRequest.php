<?php

namespace App\Http\Requests\Table;

use Illuminate\Foundation\Http\FormRequest;

class StoreTableRequest extends FormRequest
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
            "is_active" => ["required"],
            "capacity" => ["required"],
            "location" => ["required"],
            "id_thematic" => ["required"],
        ];
    }
}
