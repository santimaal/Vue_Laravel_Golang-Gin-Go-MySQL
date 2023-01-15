<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('users', function (Blueprint $table) {
            $table->id();
            $table->boolean('is_active')->default(true);
            $table->string('name');
            $table->string('password');
            $table->string('email');
            $table->string('type')->default('client');
            $table->string('img')->default('https://i.postimg.cc/W41QygPj/descarga.png');
            $table->integer('noti')->default(0);
            $table->string('chat_id')->default("");
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('users');
    }
};
