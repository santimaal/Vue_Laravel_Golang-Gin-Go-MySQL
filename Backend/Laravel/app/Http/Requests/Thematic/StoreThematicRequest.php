<?php

namespace App\Http\Requests\Thematic;

use Illuminate\Foundation\Http\FormRequest;

class StoreThematicRequest extends FormRequest
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
            "name" => ["required"],
            "location" => ["required"],
            "img" => ["required"],
        ];
    }
}
