PGDMP         $            
    x            sanpham    13.1    13.0     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16394    sanpham    DATABASE     k   CREATE DATABASE sanpham WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';
    DROP DATABASE sanpham;
                postgres    false            �            1259    16395    docs    TABLE     Y   CREATE TABLE public.docs (
    tensanpham character varying NOT NULL,
    gia integer
);
    DROP TABLE public.docs;
       public         heap    postgres    false            �          0    16395    docs 
   TABLE DATA           /   COPY public.docs (tensanpham, gia) FROM stdin;
    public          postgres    false    200           "           2606    16401    docs sanpham_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.docs
    ADD CONSTRAINT sanpham_pkey PRIMARY KEY (tensanpham);
 ;   ALTER TABLE ONLY public.docs DROP CONSTRAINT sanpham_pkey;
       public            postgres    false    200            �   (   x�KJ���4520��N��442��R9--�b���� ��     