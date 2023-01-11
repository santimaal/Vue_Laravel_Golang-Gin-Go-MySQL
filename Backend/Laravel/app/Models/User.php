<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Foundation\Auth\User as Authenticatable;
use PHPOpenSourceSaver\JWTAuth\Contracts\JWTSubject;

class User extends Authenticatable implements JWTSubject
{
    use HasFactory;
    protected $table = 'users';
    protected $fillable = ['is_active', 'name', 'password', 'email', 'type', 'img', 'noti'];
    protected $hidden = ['created_at', 'updated_at'];

    public function tables()
    {
        return $this->morphedByMany(Table::class, 'id_table');
    }

    public function getJWTIdentifier()
    {
        return $this->getKey();
    }

    public function getJWTCustomClaims()
    {
        return [];
    }
}
