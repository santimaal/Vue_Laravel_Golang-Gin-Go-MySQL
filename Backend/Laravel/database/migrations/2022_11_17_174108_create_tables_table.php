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
        Schema::create('tables', function (Blueprint $table) {
            $table->id();
            $table->boolean('is_active');
            $table->integer('capacity');
            $table->string('location');
            $table->foreignId( 'id_thematic' )->constrained( 'thematics' )->onDelete( 'cascade' );
            // $table->foreign("id_thematic")->references("id")->on("thematics");
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
        Schema::dropIfExists('tables');
    }
};
