<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Reserve extends Model
{
    use HasFactory;
    protected $table = 'reserves';
    protected $fillable = ['id_table', 'id_user', 'is_confirmed', "dateini", "datefin"];
    protected $hidden = ['created_at', 'updated_at'];
}
