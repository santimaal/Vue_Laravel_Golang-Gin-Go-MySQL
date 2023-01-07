<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Date;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{

    public function up()
    {
        Schema::create('reserves', function (Blueprint $table) {
            $table->id();
            $table->foreignId( 'id_table' )->constrained( 'tables' )->onDelete( 'cascade' );
            $table->foreignId( 'id_user' )->constrained( 'users' )->onDelete( 'cascade' );
            $table->boolean('is_confirmed')->default(false);
            $table->timestamp("dateini")->nullable();
            $table->timestamp("datefin")->nullable();
            $table->timestamps();
        });
    }

    public function down()
    {
        Schema::dropIfExists('reserves');
    }
};
