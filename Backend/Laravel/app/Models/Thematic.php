<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Thematic extends Model
{
    use HasFactory;
    protected $fillable = ['name', 'location', 'img'];
    protected $hidden = ['created_at', 'updated_at'];

    public function table()
    {
        return $this->hasMany(Table::class);
        
    }
}
