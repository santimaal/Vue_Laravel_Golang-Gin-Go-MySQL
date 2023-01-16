<?php

namespace App\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;

class UserResource extends JsonResource
{
    public function toArray($request)
    {
        return [
            'id' => $this->id,
            'is_active' => $this->is_active,
            'name' => $this->name,
            // 'password' => $this->password,
            'email' => $this->email,
            'type' => $this->type,
            'img' => $this->img,
            'noti' => $this->noti,
        ];
    }
}
