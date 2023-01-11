<?php

namespace App\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;

class TableResource extends JsonResource
{
    public function toArray($request)
    {
        return [
            'id' => $this->id,
            'is_active' => $this->is_active,
            'capacity' => $this->capacity,
            'location' => $this->location,
            'id_thematic' => $this->id_thematic,
        ];
    }
}
